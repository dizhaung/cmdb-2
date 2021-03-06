// Copyright 2020 Zhizhesihai (Beijing) Technology Limited.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/zhihu/cmdb/pkg/server"
	"github.com/zhihu/cmdb/pkg/storage/cdc"
	"github.com/zhihu/cmdb/pkg/tools/database"
	"github.com/zhihu/cmdb/pkg/tools/pd"
)

func InitServer(ctx context.Context, dsn database.DSN, pdConf *pd.Config, name cdc.DriverName, source cdc.Source) (*server.Server, error) {
	panic(wire.Build(server.Set))
}
