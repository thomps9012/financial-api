import { useAppContext } from "@/context/AppContext";

export default function ProfilePettyCashPage() {
  const { user_profile } = useAppContext();
  const { id } = user_profile;
  return (
    <main>
      <h1>Petty Cash Page for {id}</h1>
    </main>
  );
}
