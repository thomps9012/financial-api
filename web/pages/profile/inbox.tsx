import { useAppContext } from "@/context/AppContext";

function ProfileInbox() {
  const { user_profile, clearAction } = useAppContext();
  const { name, incomplete_actions } = user_profile;
  const clearNotification = (action_id: string) => {};
  return (
    <main>
      <h1>Inbox Page for {name}</h1>
      <p>{JSON.stringify(incomplete_actions, null, 2)}</p>
    </main>
  );
}

export default ProfileInbox;
