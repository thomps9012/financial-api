import { useAppContext } from "@/context/AppContext";

export default function ProfileInbox() {
  const { user_profile } = useAppContext();
  const { name, incomplete_actions } = user_profile;
  return (
    <main>
      <h1>Inbox Page for {name}</h1>
      <p>{JSON.stringify(incomplete_actions, null, 2)}</p>
    </main>
  );
}
