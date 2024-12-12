// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {menu} from '../models';
import {base} from '../models';
import {types} from '../models';
import {app} from '../models';
import {context} from '../models';
import {editors} from '../models';

export function CancelAllContexts():Promise<void>;

export function DaemonsView(arg1:menu.CallbackData):Promise<void>;

export function ExportAddress(arg1:base.Address):Promise<void>;

export function FetchAbi(arg1:number,arg2:number):Promise<types.AbiContainer>;

export function FetchAppInfo():Promise<app.AppInfo>;

export function FetchBalance(arg1:number,arg2:number):Promise<types.BalanceContainer>;

export function FetchChart(arg1:number,arg2:number):Promise<types.ChartContainer>;

export function FetchConfig(arg1:number,arg2:number):Promise<types.ConfigContainer>;

export function FetchDaemon(arg1:number,arg2:number):Promise<types.DaemonContainer>;

export function FetchIncoming(arg1:number,arg2:number):Promise<types.IncomingContainer>;

export function FetchIndex(arg1:number,arg2:number):Promise<types.IndexContainer>;

export function FetchInternal(arg1:number,arg2:number):Promise<types.InternalContainer>;

export function FetchLog(arg1:number,arg2:number):Promise<types.LogContainer>;

export function FetchManifest(arg1:number,arg2:number):Promise<types.ManifestContainer>;

export function FetchMonitor(arg1:number,arg2:number):Promise<types.MonitorContainer>;

export function FetchName(arg1:number,arg2:number):Promise<types.NameContainer>;

export function FetchNeighbor(arg1:number,arg2:number):Promise<types.NeighborContainer>;

export function FetchOutgoing(arg1:number,arg2:number):Promise<types.OutgoingContainer>;

export function FetchPin(arg1:number,arg2:number):Promise<types.PinContainer>;

export function FetchProject(arg1:number,arg2:number):Promise<types.ProjectContainer>;

export function FetchPublish(arg1:number,arg2:number):Promise<types.PublishContainer>;

export function FetchReceipt(arg1:number,arg2:number):Promise<types.ReceiptContainer>;

export function FetchSession(arg1:number,arg2:number):Promise<types.SessionContainer>;

export function FetchStatement(arg1:number,arg2:number):Promise<types.StatementContainer>;

export function FetchStatus(arg1:number,arg2:number):Promise<types.StatusContainer>;

export function FetchTrace(arg1:number,arg2:number):Promise<types.TraceContainer>;

export function FetchUpload(arg1:number,arg2:number):Promise<types.UploadContainer>;

export function FetchWizard(arg1:number,arg2:number):Promise<types.WizardContainer>;

export function FileNew(arg1:menu.CallbackData):Promise<void>;

export function FileOpen(arg1:menu.CallbackData):Promise<void>;

export function FileSave(arg1:menu.CallbackData):Promise<void>;

export function FileSaveAs(arg1:menu.CallbackData):Promise<void>;

export function Freshen():Promise<void>;

export function GetChainInfo(arg1:string):Promise<types.Chain>;

export function GetChains():Promise<Array<string>>;

export function GetConfig():Promise<types.Config>;

export function GetContext():Promise<context.Context>;

export function GetDaemon(arg1:string):Promise<string>;

export function GetDaemonOn(arg1:string):Promise<boolean>;

export function GetEnv(arg1:string):Promise<string>;

export function GetExploreUrl(arg1:string,arg2:boolean,arg3:boolean):Promise<string>;

export function GetFilter():Promise<types.Filter>;

export function GetHeaderOn(arg1:string,arg2:string):Promise<boolean>;

export function GetHistoryContainer():Promise<types.HistoryContainer>;

export function GetLastAddress():Promise<base.Address>;

export function GetLastRoute():Promise<string>;

export function GetLastTab(arg1:string):Promise<string>;

export function GetLayoutOn(arg1:string):Promise<boolean>;

export function GetMenus():Promise<menu.Menu>;

export function GetName(arg1:base.Address):Promise<string>;

export function GetState(arg1:string):Promise<string>;

export function GetTabs():Promise<Array<string>>;

export function HistoryView(arg1:menu.CallbackData):Promise<void>;

export function LoadAddress(arg1:string):Promise<void>;

export function LoadDalleImage(arg1:base.Address):Promise<boolean>;

export function LoadName(arg1:string):Promise<editors.Name>;

export function Logger(arg1:Array<string>):Promise<void>;

export function ModifyAbi(arg1:app.ModifyData):Promise<void>;

export function ModifyMonitor(arg1:app.ModifyData):Promise<void>;

export function ModifyName(arg1:app.ModifyData):Promise<void>;

export function ModifyProject(arg1:app.ModifyData):Promise<void>;

export function MonitorsView(arg1:menu.CallbackData):Promise<void>;

export function Navigate(arg1:string):Promise<void>;

export function ProjectView(arg1:menu.CallbackData):Promise<void>;

export function Reload():Promise<void>;

export function SetChain(arg1:string):Promise<void>;

export function SetDaemonOn(arg1:string,arg2:boolean):Promise<void>;

export function SetEnv(arg1:string,arg2:string):Promise<void>;

export function SetFilter(arg1:string):Promise<void>;

export function SetHeaderOn(arg1:string,arg2:string,arg3:boolean):Promise<void>;

export function SetLastAddress(arg1:string):Promise<void>;

export function SetLastRoute(arg1:string):Promise<void>;

export function SetLastTab(arg1:string,arg2:string):Promise<void>;

export function SetLayoutOn(arg1:string,arg2:boolean):Promise<void>;

export function SettingsView(arg1:menu.CallbackData):Promise<void>;

export function SharingView(arg1:menu.CallbackData):Promise<void>;

export function StepWizard(arg1:types.WizStep):Promise<types.WizState>;

export function SystemAbout(arg1:menu.CallbackData):Promise<void>;

export function SystemQuit(arg1:menu.CallbackData):Promise<void>;

export function ToggleAppFooter(arg1:menu.CallbackData):Promise<void>;

export function ToggleAppHeader(arg1:menu.CallbackData):Promise<void>;

export function ToggleAppHelp(arg1:menu.CallbackData):Promise<void>;

export function ToggleAppMenu(arg1:menu.CallbackData):Promise<void>;

export function ToggleDaemon(arg1:string):Promise<void>;

export function ToggleNextTab(arg1:menu.CallbackData):Promise<void>;

export function TogglePrevTab(arg1:menu.CallbackData):Promise<void>;

export function ToggleViewHeader(arg1:menu.CallbackData):Promise<void>;

export function UnchainedView(arg1:menu.CallbackData):Promise<void>;

export function WizardView(arg1:menu.CallbackData):Promise<void>;
