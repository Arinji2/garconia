import { END_DATE } from "./App";
import CountDownComponent from "./components/countdown";
import { VerifyHeader } from "./components/header";
import Socials from "./components/socials";

export default function Verify() {
  return (
    <>
      <VerifyHeader />
      <CountDownComponent date={END_DATE} />
      <Socials />
    </>
  );
}
