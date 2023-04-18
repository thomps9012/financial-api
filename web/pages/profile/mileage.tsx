import { useAppContext } from "@/context/AppContext";
import { Mileage_Request } from "@/types/mileage";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";

function ProfileMileagePage({
  mileage_requests,
}: {
  mileage_requests: Mileage_Request[];
}) {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Mileage Page for {name}</h1>
      <p>{JSON.stringify(mileage_requests, null, 2)}</p>
    </main>
  );
}
ProfileMileagePage.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (!credentials) {
    return {
      mileage_requests: [],
    };
  }
  const user_credentials = JSON.parse(credentials as string);
  const { data, status, statusText } = await axios.get(
    "/api/me/mileage",
    ...user_credentials
  );
  if (status != 200 || 201) {
    await axios.post("/api/error", {
      ...user_credentials,
      data: {
        user_id: user_credentials.headers.Authorization.split(" ")[1],
        error: data,
        error_path: "GET /me/mileage",
        error_message: statusText,
      },
    });
    return {
      mileage_requests: [],
    };
  }
  return {
    mileage_requests: data.data,
  };
};
export default ProfileMileagePage;
