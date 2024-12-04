import { END_DATE } from "./App";
import CountDownComponent from "./components/countdown";
import Delete from "./components/delete";
import { VerifyHeader } from "./components/header";
import Socials from "./components/socials";

export default function Verify() {
  return (
    <>
      <VerifyHeader />
      <CountDownComponent date={END_DATE} />
      <Delete />
      <Socials />
    </>
  );
}
