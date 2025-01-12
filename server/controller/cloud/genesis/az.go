/*
 * Copyright (c) 2024 Yunshan Networks
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package genesis

import (
	"time"

	"github.com/deepflowio/deepflow/server/controller/cloud/model"
	"github.com/deepflowio/deepflow/server/controller/common"
	"github.com/deepflowio/deepflow/server/libs/logger"
)

func (g *Genesis) getAZ() (model.AZ, error) {
	log.Debug("get az starting", logger.NewORGPrefix(g.orgID))
	azName := common.DEFAULT_REGION_NAME
	if g.regionLcuuid != common.DEFAULT_REGION {
		azName = g.Name
	}
	azLcuuid := common.GetUUIDByOrgID(g.orgID, azName)

	g.cloudStatsd.RefreshAPIMoniter("az", 0, time.Time{})

	az := model.AZ{
		Lcuuid:       azLcuuid,
		RegionLcuuid: g.regionLcuuid,
		Name:         azName,
	}
	g.azLcuuid = azLcuuid
	log.Debug("get az complete", logger.NewORGPrefix(g.orgID))
	return az, nil
}
