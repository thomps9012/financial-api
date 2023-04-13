import { useAppContext } from "@/context/AppContext";

export default function ProfileInbox() {
  const { user_profile } = useAppContext();
  const { id } = user_profile;
  return (
    <main>
      <h1>Inbox Page for {id}</h1>
    </main>
  );
}
