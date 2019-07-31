// Copyright (c) 2019 Baidu, Inc.
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

// load balance inside one cluster

package bfe_balance

import (
	"github.com/baidu/bfe/bfe_balance/backend"
	"github.com/baidu/bfe/bfe_basic"
	"github.com/baidu/bfe/bfe_config/bfe_cluster_conf/cluster_conf"
	"github.com/baidu/bfe/bfe_config/bfe_cluster_conf/cluster_table_conf"
	"github.com/baidu/bfe/bfe_config/bfe_cluster_conf/gslb_conf"
)

type BfeBalance interface {
	// initialize
	Init(backendConf cluster_table_conf.ClusterBackend,
		gslbBasic cluster_conf.GslbBasicConf,
		gslbConf gslb_conf.GslbClusterConf) error
	// reload config
	Reload(backendConf cluster_table_conf.ClusterBackend,
		gslbBasic cluster_conf.GslbBasicConf,
		gslbConf gslb_conf.GslbClusterConf) error
	// load balance for request
	Balance(req *bfe_basic.Request) (*backend.BfeBackend, error)
	// release
	Release()
}
