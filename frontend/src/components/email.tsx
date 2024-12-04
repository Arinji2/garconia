import { toast } from "react-hot-toast";

const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
export default function Email() {
  return (
    <form
      onSubmit={async (e) => {
        e.preventDefault();

        const email = (document.getElementById("email") as HTMLInputElement)
          .value;
        if (email.length === 0) {
          toast.error("Please Enter an Email!");
          return;
        }
        if (!emailRegex.test(email)) {
          toast.error("Invalid Email!");
          return;
        }
        const id = toast.loading("Sending Email...");
        const backendURL = import.meta.env.VITE_BACKEND_URL;
        const res = await fetch(`${backendURL}/api/add`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            email: email,
          }),
        });
        toast.dismiss(id);
        if (res.status === 200) {
          toast.success("Email Sent, Check your Inbox!");
          return;
        }
        if (res.status === 400) {
          toast.error("Email Already Exists!");
          return;
        }
        if (res.status === 409) {
          toast.error("Email Already Exists!");
          return;
        }
        toast.error("Something Went Wrong!");
      }}
      className="shrink-0 w-[80%] h-fit xl:h-[50px] gap-2 xl:gap-0 flex flex-col xl:flex-row items-center justify-center"
    >
      <input
        id="email"
        type="email"
        className=" xl:h-full text-garconia-white h-[50px] bg-garconia-offwhite/50 xl:shrink-0 w-full xl:rounded-none xl:w-[80%] outline-none rounded-md xl:rounded-l-md shadow-shadow border-[2px] border-black px-4"
        placeholder="Email..."
      />
      <button
        type="submit"
        className="px-2 bg-garconia-red w-fit xl:w-full h-[40px] xl:h-full rounded-md xl:rounded-none xl:rounded-r-md shadow-shadow border-[2px] border-black"
      >
        <p className="text-white whitespace-nowrap  md:text-2xl font-bold">
          {" "}
          NOTIFY ME!!{" "}
        </p>
      </button>
    </form>
  );
}
