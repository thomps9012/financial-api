import { useRouter } from "next/router";

export default function ServerSideError({
  request_info,
}: {
  request_info: string;
}) {
  const router = useRouter();
  return (
    <main
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
      <h3 className="REJECTED_EDIT" id="error_message">
        There was a server side error fetching your {request_info}
      </h3>
      <div className="hr-red" />
      <h3>Your error has been logged and will be resolved promptly</h3>
      <h3>
        If you would like to take a screenshot of the error and email it to App
        Support{" "}
        <a className="PENDING" href="mailto:app_support@norainc.org">
          app_support@norainc.org
        </a>{" "}
        it will expedite the issue resolution process
      </h3>
      <br />
      <a className="archive-btn" onClick={() => router.push("/")}>
        Return Home
      </a>
    </main>
  );
}
