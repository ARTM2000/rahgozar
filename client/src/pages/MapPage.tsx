import {useEffect, useRef, useState} from "react";
import 'ol/ol.css';
import {MapMenu} from "../components/MapMenu.tsx";
import MapManager from "../lib/map.ts";
import {Map} from "ol";

const MAP_NODE_ID = 'main-map';

const MapPage = () => {
    const mapRefTarget = useRef<HTMLDivElement>(null as HTMLDivElement)

    const [map, setMap] = useState<Map>(null)

    useEffect(() => {
        const currentMap = new MapManager(
            mapRefTarget.current, MapManager.getDefaultView()
        ).getMap()
        setMap(currentMap)
        return () => currentMap.setTarget(undefined)
    }, []);

    return <>
        <MapMenu/>
        <div ref={mapRefTarget} style={{position: "absolute", top: 0, left: 0, width: "100vw", height: "100vh"}}></div>
    </>
}

export default MapPage;
