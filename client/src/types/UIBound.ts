import {ActiveLayerInfoItem} from "./apiResponse.ts";

export type ActiveLayerInfoItemWithSelect = ActiveLayerInfoItem & {
    selected: boolean
}
