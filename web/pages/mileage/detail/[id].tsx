import ServerSideError from "@/components/serverSideError";
import { Mileage_Request } from "@/types/mileage";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";
import Link from "next/link";

function MileageRequestDetail({
  request_id,
  request,
}: {
  request: Mileage_Request;
  request_id: string;
}) {
  if (request_id === "") {
    return <ServerSideError request_info="Mileage Record Detail" />;
  }
  return (
    <main>
      <h1>Mileage Request Detail for {request_id}</h1>
      <Link href={`/mileage/edit/${request_id}`}>Edit Request</Link>
      <p>{JSON.stringify(request, null, 2)}</p>
    </main>
  );
}

MileageRequestDetail.getInitialProps = async (
  ctx: GetServerSidePropsContext
) => {
  const { id } = ctx.query;
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (!credentials) {
    return {
      request_id: "",
      request: {},
    };
  }
  const user_credentials = JSON.parse(credentials as string);
  const { data, status, statusText } = await axios.get("/api/mileage/detail", {
    ...user_credentials,
    data: {
      mileage_id: id,
    },
  });
  if (status != 200 || 201) {
    await axios.post("/api/error", {
      ...user_credentials,
      data: {
        user_id: user_credentials.headers.Authorization.split(" ")[1],
        error: data,
        error_path: "GET /mileage/detail",
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

export default MileageRequestDetail;
