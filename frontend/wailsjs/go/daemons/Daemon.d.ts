// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {daemons} from '../models';

export function GetState():Promise<daemons.DaemonState>;

export function Instance():Promise<daemons.Daemon>;

export function IsRunning():Promise<boolean>;

export function Pause():Promise<void>;

export function Run():Promise<void>;

export function Stop():Promise<void>;

export function String():Promise<string>;

export function Tick(arg1:Array<string>):Promise<number>;

export function Toggle():Promise<void>;
