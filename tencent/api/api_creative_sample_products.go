/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"context"
	"donson.com.cn/draining/internal/pkg/tencent/errors"
	. "donson.com.cn/draining/internal/pkg/tencent/model"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type CreativeSampleProductsApiService service

/*
CreativeSampleProductsApiService 获取创意示例商品列表
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param productCatalogId
 * @param optional nil or *CreativeSampleProductsGetOpts - Optional Parameters:
     * @param "ProductOuterIds" (optional.Interface of []string) -
     * @param "ProductSeriesId" (optional.Int64) -
     * @param "TemplateId" (optional.Int64) -
     * @param "TemplateType" (optional.String) -
     * @param "ImageId" (optional.String) -
     * @param "VideoId" (optional.String) -
     * @param "ProductFields" (optional.Interface of []string) -
     * @param "Limit" (optional.Int64) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return CreativeSampleProductsGetResponse
*/

type CreativeSampleProductsGetOpts struct {
	ProductOuterIds optional.Interface
	ProductSeriesId optional.Int64
	TemplateId      optional.Int64
	TemplateType    optional.String
	ImageId         optional.String
	VideoId         optional.String
	ProductFields   optional.Interface
	Limit           optional.Int64
	Fields          optional.Interface
}

func (a *CreativeSampleProductsApiService) Get(ctx context.Context, accountId int64, productCatalogId int64, localVarOptionals *CreativeSampleProductsGetOpts) (CreativeSampleProductsGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue CreativeSampleProductsGetResponseData
		localVarResponse    CreativeSampleProductsGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/creative_sample_products/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	localVarQueryParams.Add("product_catalog_id", parameterToString(productCatalogId, ""))
	if localVarOptionals != nil && localVarOptionals.ProductOuterIds.IsSet() {
		localVarQueryParams.Add("product_outer_ids", parameterToString(localVarOptionals.ProductOuterIds.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ProductSeriesId.IsSet() {
		localVarQueryParams.Add("product_series_id", parameterToString(localVarOptionals.ProductSeriesId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TemplateId.IsSet() {
		localVarQueryParams.Add("template_id", parameterToString(localVarOptionals.TemplateId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TemplateType.IsSet() {
		localVarQueryParams.Add("template_type", parameterToString(localVarOptionals.TemplateType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ImageId.IsSet() {
		localVarQueryParams.Add("image_id", parameterToString(localVarOptionals.ImageId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.VideoId.IsSet() {
		localVarQueryParams.Add("video_id", parameterToString(localVarOptionals.VideoId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ProductFields.IsSet() {
		localVarQueryParams.Add("product_fields", parameterToString(localVarOptionals.ProductFields.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, localVarFileKey)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, nil, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarResponse, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			if *localVarResponse.Code != 0 {
				var localVarResponseErrors []ApiErrorStruct
				if localVarResponse.Errors != nil {
					localVarResponseErrors = *localVarResponse.Errors
				}
				err = errors.NewError(localVarResponse.Code, localVarResponse.Message, localVarResponse.MessageCn, localVarResponseErrors)
				return localVarReturnValue, localVarHttpResponse.Header, err
			}
			if localVarResponse.Data == nil {
				return localVarReturnValue, localVarHttpResponse.Header, err
			} else {
				return *localVarResponse.Data, localVarHttpResponse.Header, err
			}
		} else {
			return localVarReturnValue, localVarHttpResponse.Header, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v CreativeSampleProductsGetResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse.Header, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse.Header, newErr
		}

		return localVarReturnValue, localVarHttpResponse.Header, newErr
	}

	return localVarReturnValue, localVarHttpResponse.Header, nil
}
