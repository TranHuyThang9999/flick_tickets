// Log.js
import { message } from 'antd';

export const showError = (messageText) => message.error(messageText);
export const showSuccess = (messageText) => message.success(messageText);
export const showWarning = (messageText) => message.warning(messageText);

export default function Log() {
  return null;
}