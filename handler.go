package main

import (
	"io/ioutil"
	"os"
)

/**
 *  GET要求への処理
 *
 * TODO:
 *   page要求とfavicon要求が来た際に、
 *   両方を返却してブラウザ編集が継続できるように改修
 *
 */
func getMethod(request Request) (response Response, err error) {
	htmlData, err := ioutil.ReadFile(request.Path)

	response = Response{}
	response.Status = "200"
	response.Body = htmlData
	return response, err
}

/**
 * * POST要求への処理
 *
 * TODO: 存在しないaction指定の場合は、RoutingErrorを返却させる
 * TODO: multiForm対応させる
 *
 */
func postMethod(request Request) (response Response, err error) {
	htmlData, err := ioutil.ReadFile(request.Path)

	response = Response{}
	response.Body = htmlData
	return response, err
}

/**
 * PUT要求への処理
 *
 * リソースが存在しない場合は新規で作成し、
 * リソースが存在する場合は上書き実施
 *
 */
func putMethod(request Request) (response Response, err error) {
	ioutil.WriteFile(request.Path, []byte(request.Body), 0644)

	response = Response{}
	response.Status = "204"
	return response, nil
}

/**
 * DELETE要求への処理
 *
 * リソースが存在する場合は削除する
 *
 */
func deleteMethod(request Request) (response Response, err error) {
	_, err = os.Lstat(request.Path)

	if os.IsNotExist(err) {
		msg := request.Path + " not found"
		printOut(msg, yellow, nil)

		response.Status = "202"
	} else {
		err = os.Remove(request.Path)

		response.Status = "204"
	}

	return response, err
}
