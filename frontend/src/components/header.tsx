export default function Header() {
  return (
    <div className="w-fit h-fit flex flex-col items-center justify-center gap-8">
      <p className="text-garconia-offwhite text-lg font-medium">
        Garconia Presents
      </p>
      <h1 className="font-bold text-white text-center md:text-3xl text-2xl xl:text-5xl">
        EXPLORE DEMOCRACY IN MINECRAFT
      </h1>
    </div>
  );
}
export function ConfirmHeader() {
  return (
    <div className="w-fit h-fit flex flex-col items-center justify-center gap-8">
      <p className="text-garconia-offwhite text-2xl font-medium">
        Confirm Your Email
      </p>
      <p className="text-garconia-offwhite">
        Click the button below to confirm your email
      </p>
    </div>
  );
}

export function VerifyHeader() {
  return (
    <div className="w-fit h-fit flex flex-col items-center justify-center gap-8">
      <p className="text-garconia-offwhite text-2xl font-medium">
        Thank You For Verifying Your Email
      </p>
      <p className="text-garconia-offwhite">
        We will notify you when the Server is Ready
      </p>
    </div>
  );
}
