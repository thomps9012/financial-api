import { useAppContext } from "@/context/AppContext";
import axios from "axios";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

export default function ErrorDisplay({
  message,
  path,
  error,
}: {
  error: Error;
  message: string;
  path: string;
}) {
  const [error_data, setErrorData] = useState({
    id: "",
    error_message: "",
    created_at: new Date().toISOString(),
  });
  const router = useRouter();
  const { user_profile, user_credentials } = useAppContext();
  useEffect(() => {
    async function reportError(error: Error, message: string, path: string) {
      const { data } = await axios.post("/api/error", {
        ...user_credentials,
        data: {
          user_id: user_profile.id,
          error: JSON.stringify(error),
          error_path: path,
          error_message: message,
        },
      });
      setErrorData(data.data);
    }
    reportError(error, message, path);
  }, [user_credentials, path, error, message, user_profile]);
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
        {JSON.stringify(error_data, null, 2)}
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
