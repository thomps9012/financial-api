import Link from "next/link";
import styles from "../../../styles/Home.module.css";
import { GetServerSidePropsContext } from "next";
import { getCookie } from "cookies-next";
import axios from "axios";
import { Check_Request } from "@/types/check_requests";

function CheckRequestDetail({
  request_id,
  request,
}: {
  request_id: string;
  request: Check_Request;
}) {
  const approveRequest = async (e: any) => {
    const selected_permission = (
      document.getElementById("selected_permission") as HTMLSelectElement
    ).value;
    e.preventDefault();
  };
  const rejectRequest = async (e: any) => {
    e.preventDefault();
  };
  const archiveRequest = async (e: any) => {
    e.preventDefault();
  };
  return (
    <main>
      <h1>Check Request Detail for {request_id}</h1>
      <div className="hr" />
      <Link href={`/check_request/edit/${request_id}`}>Edit Request</Link>
      <p>{JSON.stringify(request, null, 2)}</p>
    </main>
  );
}
CheckRequestDetail.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  const auth = getCookie("auth_credentials", { req: ctx.req, res: ctx.res });
  if (!auth) {
    return {
      request_id: "",
      request: {},
    };
  }
  const user_credentials = JSON.parse(auth as string);
  const { data, status, statusText } = await axios.get("/api/check/detail", {
    ...user_credentials,
    data: {
      check_request_id: id,
    },
  });
  if (status != 200 || 201) {
    await axios.post("/api/error", {
      ...user_credentials,
      data: {
        user_id: user_credentials.headers.Authorization.split(" ")[1],
        error: data,
        error_path: "GET /check/detail",
        error_message: statusText,
      },
    });
    return {
      request_id: "",
      request: {},
    };
  }
  return {
    request_id: id,
    request: data.data,
  };
};

export default CheckRequestDetail;
