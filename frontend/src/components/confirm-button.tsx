import toast from "react-hot-toast";
import { useNavigate } from "react-router-dom";

export default function ConfirmButton({
  code,
  email,
}: {
  code: string;
  email: string;
}) {
  const navigate = useNavigate();
  return (
    <button
      onClick={async () => {
        const loadingToast = toast.loading("Confirming Email...");
        const backendURL = import.meta.env.VITE_BACKEND_URL;
        const res = await fetch(`${backendURL}/api/verify`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            code: code,
            email: email,
          }),
        });
        toast.dismiss(loadingToast);
        if (res.ok) {
          toast.success("Email Confirmed!");
          navigate("/verify");
        } else {
          toast.error("Something Went Wrong!");
        }
      }}
      className="px-2 bg-garconia-red w-fit py-4 xl:w-[400px] h-[40px] xl:h-full rounded-md xl:rounded-none xl:rounded-md shadow-shadow border-[2px] border-black"
    >
      <p className="text-white whitespace-nowrap  md:text-2xl font-bold">
        CONFIRM EMAIL
      </p>
    </button>
  );
}
