import {Tile as TileLayer} from "ol/layer";
import {OSM} from "ol/source";
import {Map, View} from "ol";

const TEHRAN_CENTER_POINT = [5720067.9030910395, 4252557.027257666];

export default class MapManager {
    private readonly map: Map

    constructor(target: HTMLDivElement, view?: View | undefined) {
        this.map = new Map({
            target: target,
            layers: [this.getMapTileLayer()],
            view: view || MapManager.getDefaultView(),
            controls: []
        });
    }

    static getDefaultView(): View {
        return new View({
            center: TEHRAN_CENTER_POINT,
            zoom: 12,
            enableRotation: false,
        })
    }

    private getMapTileLayer(): TileLayer {
        /**
         * For working on dark mode: https://stackoverflow.com/a/75266899/12132998
         */
        return new TileLayer({
            preload: Infinity,
            source: new OSM(),
        })
    }

    public getMap(): Map {
        return this.map;
    }
}
