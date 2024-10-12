import {Map, View} from 'ol'
import {useEffect} from "react";
import TileLayer from "ol/layer/Tile";
import {OSM} from "ol/source";
import 'ol/ol.css';

const TEHRAN_CENTER_POINT = [5720067.9030910395, 4252557.027257666];

const MapPage = () => {
    useEffect(() => {
        const OSMLayer = new TileLayer({
            preload: Infinity,
            source: new OSM(),
        })

        const map = new Map({
            target: 'main-map',
            layers: [OSMLayer],
            view: new View({
                center: TEHRAN_CENTER_POINT,
                zoom: 12,
                enableRotation: false,
            }),
        });

        map.on("click", (e) => {
            console.log(e.map.getView().getCenter(), e.map.getView().getZoom())
        })

        return () => map.setTarget('main-map');
    }, [])

    return <div id="main-map" style={{width: "100%", height: "100vh"}}></div>
}

export default MapPage;
