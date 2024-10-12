import {Tile as TileLayer} from "ol/layer";
import {OSM} from "ol/source";
import {Map, View} from "ol";

const TEHRAN_CENTER_POINT = [5720067.9030910395, 4252557.027257666];

const getMapTileLayer = (): TileLayer => {
    /**
     * For working on dark mode: https://stackoverflow.com/a/75266899/12132998
     */
    return new TileLayer({
        preload: Infinity,
        source: new OSM(),
    })
}

export const getView = (): View => {
    return new View({
        center: TEHRAN_CENTER_POINT,
        zoom: 12,
        enableRotation: false,
    })
}

export const getMap = (target: string, view: View): Map => {
    const map = new Map({
        target: target,
        layers: [getMapTileLayer()],
        view: view,
        controls: []
    });

    map.on("click", (e) => {
        console.log({
            center: e.map.getView().getCenter(),
            zoom: e.map.getView().getZoom(),
        })
    })

    return map;
}
