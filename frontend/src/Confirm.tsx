import { useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import ConfirmButton from "./components/confirm-button";
import { ConfirmHeader } from "./components/header";
import Socials from "./components/socials";

export default function Confirm() {
  const [searchParams] = useSearchParams();
  const code = searchParams.get("code");
  const email = searchParams.get("email");
  const navigate = useNavigate();
  useEffect(() => {
    if (!code || !email) {
      navigate("/");
    }
  }, [code, email, navigate]);
  return (
    <>
      <ConfirmHeader />
      <ConfirmButton code={code!} email={email!} />
      <Socials />
    </>
  );
}
