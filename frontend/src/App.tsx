import Header from "./components/header";
import CountDownComponent from "./components/countdown";
import Email from "./components/email";
import Socials from "./components/socials";
const END_DATE = new Date("2024-12-31");
function App() {
  return (
    <>
      <div className="w-full flex flex-col items-center justify-center min-h-screen bg-garconia-background py-20 xl:h-1  ">
        <div className="w-full h-fit flex flex-col items-center gap-14 justify-start max-w-[1280px]">
          <Header />
          <CountDownComponent date={END_DATE} />
          <Email />
          <Socials />
        </div>
      </div>
    </>
  );
}

export default App;
