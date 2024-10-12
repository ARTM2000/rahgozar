import {Suspense} from 'react'
import './App.css'
import Loading from "./components/Loading.tsx";
import {Route, Routes} from "react-router-dom";
import MapPage from "./pages/MapPage.tsx";

function App() {
  return (
    <>
      <Suspense fallback={<Loading/>}>
        <Routes>
            <Route path="/map" element={<MapPage/>} />
        </Routes>
      </Suspense>
    </>
  )
}

export default App
