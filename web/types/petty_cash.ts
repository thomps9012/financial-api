import { Action } from "./actions";

export interface Petty_Cash_Request {
  id: string;
  user_id: string;
  grant_id: string;
  category: string;
  date: string;
  description: string;
  amount: number;
  receipts: string[];
  created_at: string;
  action_history: Action[];
  current_status: string;
  current_user: string;
  is_active: boolean;
}

export interface Petty_Cash_Overview {
  id: string;
  user_id: string;
  date: string;
  amount: number;
  current_status: string;
  current_user: string;
  is_active: boolean;
}

export interface Petty_Cash_Input {
  grant_id: string;
  category: string;
  date: string;
  description: string;
  amount: number;
  receipts: string[];
}

export interface Edit_Petty_Cash_Input {
  id: string;
  user_id: string;
  grant_id: string;
  category: string;
  date: string;
  description: string;
  amount: number;
  receipts: string[];
}
