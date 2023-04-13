import { useAppContext } from "@/context/AppContext";

export default function ProfilePage() {
  const { user_profile } = useAppContext();
  const { id } = user_profile;
  return (
    <main>
      <h1>Profile Page for {id}</h1>
    </main>
  );
}
