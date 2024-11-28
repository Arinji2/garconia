import Countdown from "react-countdown";

export default function CountDownComponent({ date }: { date: Date }) {
  return (
    <div className="w-fit h-fit">
      <Countdown date={date} renderer={CountDownRender} />,
    </div>
  );
  function CountDownRender({
    seconds,
    minutes,
    hours,
    days,
  }: {
    seconds: number;
    minutes: number;
    hours: number;
    days: number;
  }) {
    return (
      <section className="flex flex-row gap-9 items-center justify-center w-fit h-fit">
        <Count time={days} label="DAYS" />
        <Count time={hours} label="HOURS" />
        <Count time={minutes} label="MINUTES" />
        <Count time={seconds} label="SECONDS" />
      </section>
    );
  }

  function Count({ time, label }: { time: number; label: string }) {
    const digits = Array.from(String(time), Number);
    return (
      <div className="flex flex-row gap-6 items-center justify-center w-fit h-fit">
        <div className="w-fit h-fit flex flex-col items-center justify-center gap-3">
          <div className="w-fit h-fit flex flex-row items-center justify-center gap-3">
            <div className="bg-garconia-red w-[60px] flex flex-col items-center justify-center h-fit rounded-lg shadow-shadow border-[2px] border-black p-4">
              <p className="text-garconia-white text-2xl font-bold">
                {digits[0] ?? 0}
              </p>
            </div>
            <div className="bg-garconia-red w-[60px] flex flex-col items-center justify-center h-fit rounded-lg shadow-shadow border-[2px] border-black p-4">
              <p className="text-garconia-white text-2xl font-bold">
                {digits[1] ?? 0}
              </p>
            </div>
          </div>
          <p className="text-garconia-white text-2xl font-bold uppercase">
            {label}
          </p>
        </div>
      </div>
    );
  }
}
