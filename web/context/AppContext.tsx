import { Incomplete_Action } from "@/types/actions";
import { Grant } from "@/types/grants";
import {
  Axios_Credentials,
  Login_Input,
  User_Context_Info,
  User_Name_Info,
} from "@/types/users";
import { ReactNode, createContext, useContext, useState } from "react";
import axios from "axios";
type Props = {
  children: ReactNode;
};

type App_Context_Type = {
  user_credentials: Axios_Credentials;
  logged_in: boolean;
  user_profile: User_Context_Info;
  user_list: User_Name_Info[];
  grant_list: Grant[];
  incomplete_actions: Incomplete_Action[];
  clearAction: (action_id: string) => void;
  login: (login_info: Login_Input) => void;
  logout: () => void;
};

const default_app_context: App_Context_Type = {
  user_credentials: {
    headers: {
      Authorization: "",
    },
    withCredentials: true,
  },
  logged_in: false,
  user_profile: {
    id: "",
    admin: false,
    permissions: [""],
  },
  user_list: [],
  grant_list: [],
  incomplete_actions: [],
  clearAction: () => {},
  login: () => {},
  logout: () => {},
};

const AppContext = createContext<App_Context_Type>(default_app_context);

export function useAppContext() {
  return useContext(AppContext);
}

export function AppProvider({ children }: Props) {
  const grants = [
    {
      id: "H79TI082369",
      name: "BCORR",
    },
    {
      id: "SOR_HOUSING",
      name: "SOR Recovery Housing",
    },
    {
      id: "SOR_PEER",
      name: "SOR Peer",
    },
    {
      id: "SOR_LORAIN",
      name: "SOR Lorain 2.0",
    },
    {
      id: "H79TI085495",
      name: "RAP AID (Recover from Addition to Prevent Aids)",
    },
    {
      id: "2020-JY-FX-0014",
      name: "JSBT (OJJDP) - Jumpstart For A Better Tomorrow",
    },
    {
      id: "H79SP082264",
      name: "HIV Navigator",
    },
    {
      id: "H79SP082475",
      name: "SPF (HOPE 1)",
    },
    {
      id: "SOR_TWR",
      name: "SOR 2.0 - Together We Rise",
    },
    {
      id: "H79TI083370",
      name: "BSW (Bridge to Success Workforce)",
    },
    {
      id: "H79SM085150",
      name: "CCBHC",
    },
    {
      id: "H79TI083662",
      name: "IOP New Syrenity Intensive outpatient Program",
    },
    {
      id: "TANF",
      name: "TANF",
    },
    {
      id: "H79SP081048",
      name: "STOP Grant",
    },
    {
      id: "H79TI085410",
      name: "N MAT (NORA Medication-Assisted Treatment Program)",
    },
  ];
  const user_list = [{ id: "", name: "", active: false }];
  const [user_token, setToken] = useState("");
  const [user_logged_in, setLoggedIn] = useState(false);
  const [user_info, setUserInfo] = useState({
    id: "",
    admin: false,
    permissions: [""],
  });
  const [incompleteActions, setIncompleteActions] = useState(
    new Array<Incomplete_Action>()
  );
  const clearAction = (action_id: string) => {
    const new_state = incompleteActions.filter(({ id }) => id != action_id);
    setIncompleteActions(new_state);
  };
  const login_user = async (user_info: Login_Input) => {
    try {
      const res = await axios.post("/api/auth/login", user_info);
      console.log(res);
      if ((res && res.data?.data && res.status === 200) || res.status === 201) {
        setLoggedIn(true);
        setToken(res.data.data.token);
        setUserInfo({
          id: res.data.data.user_id,
          admin: res.data.data.admin,
          permissions: res.data.permissions,
        });
      }
    } catch (error) {
      console.error(error);
    }
  };
  const logout_user = async () => {
    try {
      const res = await axios.post("/api/auth/logout");
      if (res.status === 200) {
        setLoggedIn(false);
        setToken("");
        setUserInfo({
          id: "",
          admin: false,
          permissions: [""],
        });
      }
    } catch (error) {
      console.error(error);
    }
  };
  const value = {
    user_profile: user_info,
    user_credentials: {
      headers: {
        Authorization: `Bearer ${user_token}`,
      },
      withCredentials: true,
    },
    logged_in: user_logged_in,
    user_list: user_list,
    grant_list: grants,
    incomplete_actions: incompleteActions,
    clearAction: clearAction,
    login: login_user,
    logout: logout_user,
  };

  return <AppContext.Provider value={value}>{children}</AppContext.Provider>;
}
