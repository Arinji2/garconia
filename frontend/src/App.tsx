import { Toaster } from "react-hot-toast";
import { Route, Routes } from "react-router-dom";
import Confirm from "./Confirm";
import Home from "./Home";
import Verify from "./Verify";
function App() {
  return (
    <>
      <div>
        <Toaster
          position="top-right"
          toastOptions={{
            style: {
              borderRadius: "10px",
              background: "#333",
              color: "#fff",
            },
          }}
        />
      </div>

      <div className="w-full flex flex-col items-center justify-start md:justify-center min-h-screen bg-garconia-background md:py-10 py-5 xl:py-20 xl:h-1  ">
        <img
          src="/verify-bg.webp"
          alt="Verify Image"
          className="object-cover w-full h-[100svh] fixed top-0 left-0 "
        />
        <div className="w-full h-[100svh] fixed top-0 left-0 z-10 bg-black/70"></div>
        <div className="w-full h-fit flex flex-col items-center gap-14 justify-start max-w-[1280px] z-20">
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/verify" element={<Verify />} />
            <Route path="/confirm" element={<Confirm />} />
          </Routes>
        </div>
      </div>
    </>
  );
}

export default App;
