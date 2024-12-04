import { END_DATE } from "./App";
import CountDownComponent from "./components/countdown";
import Email from "./components/email";
import Header from "./components/header";
import Socials from "./components/socials";

export default function Home() {
  return (
    <>
      <Header />
      <CountDownComponent date={END_DATE} />
      <Email />
      <Socials />
    </>
  );
}
