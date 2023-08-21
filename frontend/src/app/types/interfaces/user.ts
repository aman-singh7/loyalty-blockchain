import { AccountType } from '../enums/contractEnums';

export interface User {
  uid: string;
  platformUid: string;
  accountType: AccountType;
  name: string;
  walletId: string;
}
