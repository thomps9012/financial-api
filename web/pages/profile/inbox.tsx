import { useAppContext } from "@/context/AppContext";

export default function ProfileInbox() {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Inbox Page for {name}</h1>
    </main>
  );
}
