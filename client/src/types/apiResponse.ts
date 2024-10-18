export type ServerResponse<T = any> = {
    track_id: string,
    error: boolean,
    message: string,
    data: T,
}

export type ActiveLayerInfoItem = {
    id: number,
    name: string,
    title: string,
    image: string,
}

export type AllActiveMapLayersResponse = {
    layers: ActiveLayerInfoItem[],
};
