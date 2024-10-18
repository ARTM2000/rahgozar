import {httpClient} from "../lib/httpClient.ts";
import {AllActiveMapLayersResponse, ServerResponse} from "../types/apiResponse.ts";
import {AxiosResponse} from "axios";

export const getAllActiveMapLayers = async (): Promise<AllActiveMapLayersResponse> => {
    const response =
        await httpClient.get<any, AxiosResponse<ServerResponse<AllActiveMapLayersResponse>>>("/map-layers/v1/layers-list/");
    return response.data.data;
}
