import { useAppContext } from "@/context/AppContext";

function ProfilePage() {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Profile Page for {name}</h1>
      <p>Info</p>
      {JSON.stringify(user_profile, null, 2)}
    </main>
  );
}
export default ProfilePage;
