// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
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

package oam

import (
        "bytes"
        "io/ioutil"
        "net/http"
        "net/http/httptest"
        "github.com/gorilla/mux"
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const testdataBasepath = "../../test/oam/ngc-apistub-testdata/"
const postdataBasepath = "../../test/oam/cnca-cli-scripts/json/"

func TestOam(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Oam Suite")
}

var _ = Describe("NGC_Proxy", func() {

        Describe("Proxy init", func() {
           It("Will init proxy",
           func() {
           testPath := testdataBasepath + "testdata_00.json"
           Expect(InitProxy("v", "Flexcore", testPath)).NotTo(BeNil())
           Expect(InitProxy("v", "APISTUB", testPath)).To(BeNil())
           })
        })

        Describe("Proxy GetlAll", func() {
           It("Will use proxy to GetAll",
           func() {
           testPath := testdataBasepath + "testdata_00.json"
           Expect(InitProxy("valid", "Flexcore", testPath)).NotTo(BeNil())
           req, err := http.NewRequest("GET", "/services", nil)
           Expect(err).ShouldNot(HaveOccurred())
           rsp := httptest.NewRecorder()
           ProxyGetAll(rsp,req)
           Expect(rsp.Code).To(Equal(http.StatusNotFound))

           })
        })

        Describe("Proxy Get", func() {
           It("Will use proxy to Get",
           func() {
           testPath := testdataBasepath + "testdata_00.json"
           Expect(InitProxy("valid", "Flexcore", testPath)).NotTo(BeNil())
           req, err := http.NewRequest("GET", "/services", nil)
           Expect(err).ShouldNot(HaveOccurred())
           rsp := httptest.NewRecorder()
           ProxyGet(rsp,req)
           Expect(rsp.Code).NotTo(Equal(http.StatusOK))

           })
        })

        Describe("Proxy Del", func() {
           It("Will use proxy to Del",
           func() {
           testPath := testdataBasepath + "testdata_00.json"
           Expect(InitProxy("valid", "Flexcore", testPath)).NotTo(BeNil())
           req, err := http.NewRequest("DEL", "/services/1", nil)
           Expect(err).ShouldNot(HaveOccurred())
           rsp := httptest.NewRecorder()
           ProxyDel(rsp,req)
           Expect(rsp.Code).NotTo(Equal(http.StatusOK))

           })
        })

        Describe("Proxy Add", func() {
           It("Will use proxy to Add",
           func() {
           testPath := testdataBasepath + "testdata_00.json"
           Expect(InitProxy("valid", "Flexcore", testPath)).NotTo(BeNil())
           req, err := http.NewRequest("POST", "/services/1", nil)
           Expect(err).ShouldNot(HaveOccurred())
           rsp := httptest.NewRecorder()
           ProxyAdd(rsp,req)
           Expect(rsp.Code).NotTo(Equal(http.StatusOK))

           })
        })

        Describe("Proxy Update", func() {
           It("Will use proxy to Update",
           func() {
           testPath := testdataBasepath + "testdata_00.json"
           Expect(InitProxy("valid", "Flexcore", testPath)).NotTo(BeNil())
           req, err := http.NewRequest("PATCH", "/services/1", nil)
           Expect(err).ShouldNot(HaveOccurred())
           rsp := httptest.NewRecorder()
           ProxyUpdate(rsp,req)
           Expect(rsp.Code).NotTo(Equal(http.StatusOK))

           })
        })

})

var _ = Describe("NGC_APIStub", func() {
        BeforeEach(func() {
           APIStubReset()
        })


        AfterEach(func() {
           APIStubReset()
        })

        Describe("APISTUB init", func() {
           It("Will init APSTUB",
           func() {
           Expect(APIStubInit("nonexistent-file")).NotTo(BeNil())
           Expect(APIStubInit("conf")).NotTo(BeNil())
           Expect(APIStubInit(testdataBasepath+"testdata_00.json")).To(BeNil())
           Expect(len(AllRecords)).To(Equal(0))
           Expect(NewRecordAFServiceID).To(Equal(AFServiceIDBaseValue))
           Expect(APIStubInit(testdataBasepath+"testdata_01.json")).To(BeNil())
           Expect(len(AllRecords)).To(Equal(1))
           Expect(NewRecordAFServiceID).To(Equal(AFServiceIDBaseValue+1))
           })
        })

        Describe("APISTUB reset", func() {
           It("Will reset APSTUB",
           func() {
           Expect(APIStubReset()).To(BeNil())
           })
        })

        Describe("APISTUB Add", func() {
           It("Will Add new Record",
           func() {
           reqBody, err := ioutil.ReadFile(postdataBasepath + "POST001.json") 
           Expect(err).ShouldNot(HaveOccurred())
           reqBodyBytes := bytes.NewReader(reqBody)
           req, _ := http.NewRequest(http.MethodPost,"/services",reqBodyBytes)
           rsp := httptest.NewRecorder()
           expected := "{\"afServiceId\":\"123457\"}"
           APIStubAdd(rsp, req)
           Expect(rsp.Code).To(Equal(http.StatusOK))
           Expect(rsp.Body.String()).To(Equal(expected))

           reqBody, err = ioutil.ReadFile(postdataBasepath + "POST002.json") 
           Expect(err).ShouldNot(HaveOccurred())
           reqBodyBytes = bytes.NewReader(reqBody)
           req, _ = http.NewRequest(http.MethodPost,"/services",reqBodyBytes)
           rsp = httptest.NewRecorder()
           expected = "{\"afServiceId\":\"123458\"}"
           APIStubAdd(rsp, req)
           Expect(rsp.Code).To(Equal(http.StatusOK))
           Expect(rsp.Body.String()).To(Equal(expected))

           reqBody, err = ioutil.ReadFile(postdataBasepath + "POST003.json") 
           Expect(err).ShouldNot(HaveOccurred())
           reqBodyBytes = bytes.NewReader(reqBody)
           req, _ = http.NewRequest(http.MethodPost,"/services",reqBodyBytes)
           rsp = httptest.NewRecorder()
           expected = "{\"afServiceId\":\"123459\"}"
           APIStubAdd(rsp, req)
           Expect(rsp.Code).To(Equal(http.StatusOK))
           Expect(rsp.Body.String()).To(Equal(expected))
                               
           })
        })

        Describe("APISTUB Update", func() {
           It("Will Update Record",
           func() {
           APIStubInit(testdataBasepath + "testdata_01.json")
           reqBody, err := ioutil.ReadFile(postdataBasepath + "POST001.json") 
           Expect(err).ShouldNot(HaveOccurred())
           reqBodyBytes := bytes.NewReader(reqBody)
           req, _ := http.NewRequest("PATCH","/services/1",reqBodyBytes)
           vars := map[string]string{
           "afServiceId": "123457",
           }                               
           req = mux.SetURLVars(req, vars)
           rsp := httptest.NewRecorder()
           APIStubUpdate(rsp, req)
           Expect(rsp.Code).To(Equal(http.StatusOK))
                                
           req, _ = http.NewRequest("PATCH","/services/2",nil)
           vars = map[string]string{
           "afServiceId": "123458",
           }                               
           req = mux.SetURLVars(req, vars)
           rsp = httptest.NewRecorder()
           APIStubDel(rsp, req)
           Expect(rsp.Code).NotTo(Equal(http.StatusOK))
                               
           })
       })
 
       Describe("APISTUB Del", func() {
           It("Will Delete Record",
           func() {
           APIStubInit(testdataBasepath + "testdata_01.json")
           req, _ := http.NewRequest("DELETE","/services/123457",nil)
           vars := map[string]string{
           "afServiceId": "123457",
           }                               
           req = mux.SetURLVars(req, vars)
           rsp := httptest.NewRecorder()
           APIStubDel(rsp, req)
           Expect(rsp.Code).To(Equal(http.StatusOK))
           Expect(len(AllRecords)).To(Equal(0))
           
           req, _ = http.NewRequest("DELETE","/services/123457",nil)
           vars = map[string]string{
           "afServiceId": "123457",
           }                               
           req = mux.SetURLVars(req, vars)
           rsp = httptest.NewRecorder()
           APIStubDel(rsp, req)
           Expect(rsp.Code).NotTo(Equal(http.StatusOK))
                               
           })
       })

       Describe("APISTUB Get", func() {
           It("Will Get one Record",
           func() {
           APIStubInit(testdataBasepath + "testdata_01.json")
           req, _ := http.NewRequest("GET","/services/123457",nil)
           vars := map[string]string{
           "afServiceId": "123457",
           }                               
           req = mux.SetURLVars(req, vars)
           rsp := httptest.NewRecorder()
           APIStubGet(rsp, req)
           Expect(rsp.Code).To(Equal(http.StatusOK))


           req, _ = http.NewRequest("GET","/services/123458",nil)
           vars = map[string]string{
           "afServiceId": "123458",
           }                               
           req = mux.SetURLVars(req, vars)
           rsp = httptest.NewRecorder()
           APIStubGet(rsp, req)
           Expect(rsp.Code).NotTo(Equal(http.StatusOK))

                               
           })
       })

       Describe("APISTUB Getll", func() {
           It("Will GetAll Records",
           func() {
           APIStubInit(testdataBasepath + "testdata_01.json")
           req, err := http.NewRequest("GET", "/services", nil)
           Expect(err).ShouldNot(HaveOccurred())
           rsp := httptest.NewRecorder()
           APIStubGetAll(rsp, req)
           Expect(rsp.Code).To(Equal(http.StatusOK))
                               
          })
        })
})

