package utils

import (
	"DnspodCut/structs"
	"encoding/json"
	"fmt"
	"log"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

func FindDns(config structs.Config, dns structs.Dns) ([]structs.RecordList, error) {
	credential := common.NewCredential(
		config.SecretId,
		config.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewDescribeRecordListRequest()

	request.Domain = common.StringPtr("guiyunweb.com")
	request.Subdomain = common.StringPtr(dns.SubDomain)
	request.RecordType = common.StringPtr(dns.RecordType)

	// 返回的resp是一个DescribeRecordListResponse的实例，与请求对象对应
	response, err := client.DescribeRecordList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		log.Printf("返回API错误: %s", err)
		return []structs.RecordList{}, err
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())

	data, err := jsonToStruct(response.ToJsonString())
	if err != nil {
		log.Printf("转换错误 : %s", err)
	}
	return data.Response.RecordList, nil
}

func UpdateDns(config structs.Config, dns structs.Dns, ip structs.RecordList, status string) {
	credential := common.NewCredential(
		config.SecretId,
		config.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dnspod.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewModifyRecordRequest()

	request.Domain = common.StringPtr(dns.Domain)
	request.RecordType = common.StringPtr(ip.Type)
	request.RecordLine = common.StringPtr(ip.Line)
	request.Value = common.StringPtr(ip.Value)
	request.RecordId = common.Uint64Ptr(ip.RecordId)
	request.SubDomain = common.StringPtr(dns.SubDomain)
	request.Status = common.StringPtr(status)

	// 返回的resp是一个ModifyRecordResponse的实例，与请求对象对应
	response, err := client.ModifyRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("返回API错误: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}

func jsonToStruct(jsonStr string) (structs.Response, error) {
	var person structs.Response
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		return structs.Response{}, err
	}
	return person, nil
}
