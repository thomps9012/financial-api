import { useAppContext } from "@/context/AppContext";
import { Check_Request } from "@/types/check_requests";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";

function ProfileCheckPage({
  check_requests,
}: {
  check_requests: Check_Request[];
}) {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Check Request Page for {name}</h1>
      <p>{JSON.stringify(check_requests, null, 2)}</p>
    </main>
  );
}
ProfileCheckPage.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (!credentials) {
    return {
      check_requests: [],
    };
  }
  const user_credentials = JSON.parse(credentials as string);
  const { data, status, statusText } = await axios.get(
    "/api/me/check",
    ...user_credentials
  );
  if (status != 200 || 201) {
    await axios.post("/api/error", {
      ...user_credentials,
      data: {
        user_id: user_credentials.headers.Authorization.split(" ")[1],
        error: data,
        error_path: "GET /me/check",
        error_message: statusText,
      },
    });
    return {
      check_requests: [],
    };
  }
  return {
    check_requests: data.data,
  };
};
export default ProfileCheckPage;
