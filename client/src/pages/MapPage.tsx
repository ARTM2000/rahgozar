import {useEffect, useState} from "react";
import 'ol/ol.css';
import {getMap, getView} from "../lib/map/base.ts";

const MAP_NODE_ID = 'main-map';

const MapPage = () => {
    const [map, setMap] = useState<Map<any, any>>(null)

    useEffect(() => {
        const currentMap = getMap(MAP_NODE_ID, getView())
        setMap(currentMap as Map<any, any>)
    }, []);

    return <div id={MAP_NODE_ID} style={{width: "100%", height: "100vh"}}></div>
}

export default MapPage;
