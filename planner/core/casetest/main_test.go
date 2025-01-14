// Copyright 2023 PingCAP, Inc.
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

package casetest

import (
	"flag"
	"testing"

	"github.com/pingcap/tidb/testkit/testdata"
	"github.com/pingcap/tidb/testkit/testmain"
	"github.com/pingcap/tidb/testkit/testsetup"
	"go.uber.org/goleak"
)

var testDataMap = make(testdata.BookKeeper)

func TestMain(m *testing.M) {
	testsetup.SetupForCommonTest()

	flag.Parse()

	testDataMap.LoadTestSuiteData("testdata", "integration_partition_suite")
	testDataMap.LoadTestSuiteData("testdata", "plan_normalized_suite")
	testDataMap.LoadTestSuiteData("testdata", "stats_suite")

	testDataMap.LoadTestSuiteData("testdata", "point_get_plan")
	testDataMap.LoadTestSuiteData("testdata", "expression_rewriter_suite")
	testDataMap.LoadTestSuiteData("testdata", "partition_pruner")
	testDataMap.LoadTestSuiteData("testdata", "integration_suite")
	testDataMap.LoadTestSuiteData("testdata", "window_push_down_suite")

	testDataMap.LoadTestSuiteData("testdata", "json_plan_suite")

	testDataMap.LoadTestSuiteData("testdata", "predicate_simplification")

	opts := []goleak.Option{
		goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"),
		goleak.IgnoreTopFunction("github.com/lestrrat-go/httprc.runFetchWorker"),
		goleak.IgnoreTopFunction("go.etcd.io/etcd/client/pkg/v3/logutil.(*MergeLogger).outputLoop"),
		goleak.IgnoreTopFunction("gopkg.in/natefinch/lumberjack%2ev2.(*Logger).millRun"),
		goleak.IgnoreTopFunction("github.com/tikv/client-go/v2/txnkv/transaction.keepAlive"),
		goleak.IgnoreTopFunction("go.opencensus.io/stats/view.(*worker).start"),
	}

	callback := func(i int) int {
		testDataMap.GenerateOutputIfNeeded()
		return i
	}

	goleak.VerifyTestMain(testmain.WrapTestingM(m, callback), opts...)
}

func GetIntegrationPartitionSuiteData() testdata.TestData {
	return testDataMap["integration_partition_suite"]
}

func GetPlanNormalizedSuiteData() testdata.TestData {
	return testDataMap["plan_normalized_suite"]
}

func GetStatsSuiteData() testdata.TestData {
	return testDataMap["stats_suite"]
}

func GetPointGetPlanData() testdata.TestData {
	return testDataMap["point_get_plan"]
}

func GetExpressionRewriterSuiteData() testdata.TestData {
	return testDataMap["expression_rewriter_suite"]
}

func GetPartitionPrunerData() testdata.TestData {
	return testDataMap["partition_pruner"]
}

func GetIntegrationSuiteData() testdata.TestData {
	return testDataMap["integration_suite"]
}

func GetWindowPushDownSuiteData() testdata.TestData {
	return testDataMap["window_push_down_suite"]
}

func GetJSONPlanSuiteData() testdata.TestData {
	return testDataMap["json_plan_suite"]
}

func GetPredicateSimplificationTestData() testdata.TestData {
	return testDataMap["predicate_simplification"]
}
