// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	configadmission "github.com/onosproject/onos-operator/pkg/admission/config"
	configapi "github.com/onosproject/onos-operator/pkg/apis/config"
	configctrl "github.com/onosproject/onos-operator/pkg/controller/config"
	"github.com/onosproject/onos-operator/pkg/controller/util/leader"
	"github.com/onosproject/onos-operator/pkg/controller/util/ready"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"os"
	"runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

var log = logging.GetLogger("config-operator")

func printVersion() {
	log.Info(fmt.Sprintf("Go Version: %s", runtime.Version()))
	log.Info(fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH))
}

func main() {
	var namespace string
	if len(os.Args) > 1 {
		namespace = os.Args[1]
	}

	printVersion()

	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Become the leader before proceeding
	_ = leader.Become(context.TODO())

	r := ready.NewFileReady()
	err = r.Set()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer func() {
		_ = r.Unset()
	}()

	// Create a new Cmd to provide shared dependencies and start components
	mgr, err := manager.New(cfg, manager.Options{Namespace: namespace})
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	mgr.GetClient()

	log.Info("Registering components")

	// Setup Scheme for all resources
	if err := configapi.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Setup all Controllers
	if err := configctrl.AddControllers(mgr); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Setup all webhooks
	if err := configadmission.RegisterWebhooks(mgr); err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Info("Starting the operator")

	// Start the Cmd
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		log.Error(err, "operator exited non-zero")
		os.Exit(1)
	}
}
