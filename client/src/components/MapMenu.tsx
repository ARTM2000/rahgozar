import {useEffect, useState} from "react";
import {getAllActiveMapLayers} from "../api/mapApis.ts";
import {ActiveLayerInfoItem} from "../types/apiResponse.ts";
import {ActiveLayerInfoItemWithSelect} from "../types/UIBound.ts";

export const MapMenu = () => {
    const [layersList, setLayersList] = useState<ActiveLayerInfoItemWithSelect[]>([])

    const handleLayerItemClick = (l: ActiveLayerInfoItem) => {
        const newLayersList = layersList.map(layer => layer.id === l.id ? {...layer, selected: !layer.selected} : layer)
        setLayersList(newLayersList)
    }

    useEffect(() => {
        getAllActiveMapLayers().then(layersData => {
            const formattedLayersData = layersData.layers.map(l => ({...l, selected: false}))
            setLayersList(formattedLayersData);
            console.log("layers list info received!")
        }).catch(err => console.error(err))
    }, []);

    return <div className={"z-1 position-absolute col-12 col-sm-8 col-md-6 col-xl-4 bg-transparent"} style={{
        right: 0,
        top: 0,
        maxWidth: "500px",
        height: "auto",
    }}>
        <div className={"text-body bg-body position-relative m-3 pb-2 rounded-3 h-75"}>
            <p className={'px-3 pt-3 text-center mb-3'}>نوع حمل و نقل مورد نظر را انتخاب کنید:</p>
            <div className={"row flex-row-reverse my-3 mx-2"} style={{minWidth: 0}}>
                {layersList.map(l =>
                    <div
                        key={l.id}
                        className={'col-4 map-layer-item mb-3'}
                        style={{textAlign: 'center'}}
                        onClick={() => handleLayerItemClick(l)}
                    >
                        <div className={`border border-0 rounded-3 p-3 text-center ${l.selected ? 'text-light bg-primary' : 'bg-secondary'}`}>
                            {l.title}
                        </div>
                    </div>)}
            </div>
        </div>
    </div>
}
