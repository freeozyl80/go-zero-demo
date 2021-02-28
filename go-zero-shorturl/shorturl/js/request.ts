import webapi from "axios"
import * as components from "./requestComponents"
export * from "./requestComponents"

/**
 * @description 
 * @param params
 */
export function api(params: components.GreetReqParams) {
	return webapi.get<components.GreetResp>("/shorten", params)
}
