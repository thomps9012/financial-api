import { useRouter } from "next/router";

export default function ErrorDisplay({
  message,
  path,
}: {
  message: string;
  path: string;
}) {
  const router = useRouter();
  return (
    <div
      style={{
        margin: 100,
        padding: 20,
        display: "flex",
        flexDirection: "column",
        flexWrap: "wrap",
      }}
    >
      <h1 className="REJECTED">Error!</h1>
      <div className="hr-red" />
      <h3 className="REJECTED_EDIT">{message}</h3>
      <h3 className="REJECTED_EDIT">Occurred on {path}</h3>
      <div className="hr-red" />
      {/* add in report api call */}
      <h3>
        Please take a screenshot of this screen and email it to App Support{" "}
        <a className="PENDING" href="mailto:app_support@norainc.org">
          app_support@norainc.org
        </a>{" "}
        for expedited issue resolution
      </h3>
      <br />
      <a className="archive-btn" onClick={() => router.push("/")}>
        Return Home
      </a>
    </div>
  );
}
