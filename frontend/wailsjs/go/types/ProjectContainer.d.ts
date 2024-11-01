// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {base} from '../models';
import {types} from '../models';

export function Clean(arg1:base.Address):Promise<Array<base.Address>|base.Address>;

export function ForEveryHistory(arg1:types.EveryAddressFn,arg2:any):Promise<boolean>;

export function Load(arg1:string):Promise<types.ProjectFile>;

export function NeedsUpdate(arg1:boolean):Promise<boolean>;

export function Save(arg1:string,arg2:base.Address):Promise<void>;

export function ShallowCopy():Promise<types.Containerer>;

export function String():Promise<string>;

export function Summarize():Promise<void>;
