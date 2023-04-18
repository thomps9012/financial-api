import ServerSideError from "@/components/serverSideError";
import { useAppContext } from "@/context/AppContext";

function ProfilePage() {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  if (user_profile.id === "") {
    return <ServerSideError request_info="Personal Profile Page" />;
  }
  return (
    <main>
      <h1>Profile Page for {name}</h1>
      <p>Info</p>
      {JSON.stringify(user_profile, null, 2)}
    </main>
  );
}
export default ProfilePage;
