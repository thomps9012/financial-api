import { Check_Request_Overview } from "@/types/check_requests";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";
import Link from "next/link";

function UserCheckPage({
  user_id,
  user_name,
  requests,
}: {
  user_id: string;
  user_name: string;
  requests: Check_Request_Overview[];
}) {
  return (
    <main>
      <h1>Check Requests for {user_name}</h1>
      <p>{JSON.stringify(requests, null, 2)}</p>
      <Link href={`/${user_id}`}>
        &larr; · &#8594; · &#x2192; Back to {user_name} Overview
      </Link>
    </main>
  );
}

UserCheckPage.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (!credentials) {
    return {
      user_id: "",
      user_name: "",
      requests: [],
    };
  }
  const auth = JSON.parse(credentials as string);
  const { data, status, statusText } = await axios.get("/api/user/check", {
    ...auth,
    data: {
      user_id: id,
    },
  });
  if (status != 200 || 201) {
    await axios.post("/api/error", {
      ...auth,
      data: {
        user_id: auth.headers.Authorization.split(" ")[1],
        error: data,
        error_path: "GET /user/check",
        error_message: statusText,
      },
    });
    return {
      user_id: "",
      user_name: "",
      requests: [],
    };
  }
  const {
    data: user_data,
    status: user_status,
    statusText: user_statusText,
  } = await axios.get("/api/user/name", {
    ...auth,
    data: {
      user_id: id,
    },
  });
  if (user_status != 200 || 201) {
    await axios.post("/api/error", {
      ...auth,
      data: {
        user_id: auth.headers.Authorization.split(" ")[1],
        error: user_data,
        error_path: "GET /user/check",
        error_message: user_statusText,
      },
    });
    return {
      user_id: "",
      user_name: "",
      requests: [],
    };
  }
  return {
    user_id: id,
    user_name: user_data.data,
    requests: data.data,
  };
};

export default UserCheckPage;
