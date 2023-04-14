import { useAppContext } from "@/context/AppContext";

export default function ProfilePettyCashPage() {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Petty Cash Page for {name}</h1>
    </main>
  );
}
