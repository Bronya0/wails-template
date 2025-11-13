const TimeUtil = {
    /**
     * 获取当前 Date 对象
     * @returns {Date} 当前 Date 对象
     */
    getCurrentDate: function () {
        return new Date();
    },

    // 年月日的字符串，不包括 时间部分
    getCurrentDateStr: function () {
        const date = new Date();
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0'); // 补零
        const day = String(date.getDate()).padStart(2, '0'); // 补零
        return `${year}-${month}-${day}`;
    },

    // 年月日时间字符串，无连接符号
    getCurrentDateTimeCleanStr: function () {
        const date = new Date();
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0'); // 补零
        const day = String(date.getDate()).padStart(2, '0'); // 补零
        const hours = String(date.getHours()).padStart(2, '0');
        const minutes = String(date.getMinutes()).padStart(2, '0');
        const seconds = String(date.getSeconds()).padStart(2, '0');
        return `${year}${month}${day}${hours}${minutes}${seconds}`;
    },

    getCurrentDateTimeStr: function () {
        return new Date().toLocaleString();
    },
    /**
     * 获取当前秒级时间戳
     * @returns {number} 当前秒级时间戳
     */
    getCurrentTimestampSeconds: function () {
        return Math.floor(Date.now() / 1000);
    },

    /**
     * 获取当前毫秒级时间戳
     * @returns {number} 当前毫秒级时间戳
     */
    getCurrentTimestampMilliseconds: function () {
        return Date.now();
    },

    /**
     * 获取当前格式化的日期时间字符串 (yyyy-MM-dd HH:mm:ss)
     * @returns {string} 格式化的日期时间字符串
     */
    getCurrentFormattedDateTime: function () {
        return this.dateToFormattedDateTime(new Date());
    },
    /**
     * 获取今天24点的时间戳
     */
    getToday24Timestamp() {
        const now = new Date();
        const year = now.getFullYear();
        const month = String(now.getMonth() + 1).padStart(2, '0');
        const day = String(now.getDate()).padStart(2, '0');
        const formattedDate = `${year}-${month}-${day}`;
        const date = new Date(`${formattedDate}T23:59:59`);
        return Math.floor(date.getTime() / 1000);
    },
    // 获取当前时间加指定小时后的时间
    getFutureHourTimestamp(hours) {
        const now = new Date();
        const future = new Date(now.getTime() + hours * 60 * 60 * 1000);
        return Math.floor(future.getTime() / 1000);
    },
    // 当前减指定天数的0点
    getDayStartTimestampMilliseconds(n) {
        const now = new Date(); // 当前时间
        const yesterday = new Date(now); // 克隆当前时间
        yesterday.setDate(now.getDate() - n); // 设置为昨天
        yesterday.setHours(0, 0, 0, 0); // 设置时间为当天的 0 点 0 分 0 秒 0 毫秒
        return yesterday.getTime(); // 返回时间戳（毫秒）
    },
    /**
     * Date 对象转换为秒级时间戳
     * @param {Date} date Date 对象
     * @returns {number} 秒级时间戳
     */
    dateToTimestampSeconds: function (date) {
        return Math.floor(date.getTime() / 1000);
    },

    /**
     * Date 对象转换为毫秒级时间戳
     * @param {Date} date Date 对象
     * @returns {number} 毫秒级时间戳
     */
    dateToTimestampMilliseconds: function (date) {
        return date.getTime();
    },

    /**
     * Date 对象转换为格式化的日期时间字符串 (yyyy-MM-dd HH:mm:ss)
     * @param {Date} date Date 对象
     * @returns {string} 格式化的日期时间字符串
     */
    dateToFormattedDateTime: function (date) {
        const year = date.getFullYear();
        const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从 0 开始
        const day = String(date.getDate()).padStart(2, '0');
        const hour = String(date.getHours()).padStart(2, '0');
        const minute = String(date.getMinutes()).padStart(2, '0');
        const second = String(date.getSeconds()).padStart(2, '0');
        return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
    },

    /**
     * 秒级时间戳转换为 Date 对象
     * @param {number} timestamp 秒级时间戳
     * @returns {Date} Date 对象
     */
    timestampSecondsToDate: function (timestamp) {
        return new Date(timestamp * 1000);
    },

    /**
     * 毫秒级时间戳转换为 Date 对象
     * @param {number} timestamp 毫秒级时间戳
     * @returns {Date} Date 对象
     */
    timestampMillisecondsToDate: function (timestamp) {
        return new Date(timestamp);
    },

    /**
     * 格式化的日期时间字符串 (yyyy-MM-dd HH:mm:ss) 转换为 Date 对象
     * @param {string} formattedDateTime 格式化的日期时间字符串
     * @returns {Date} Date 对象
     */
    formattedStringToDate: function (formattedDateTime) {
        return new Date(formattedDateTime); // Date 构造函数可以直接解析 yyyy-MM-dd HH:mm:ss 格式
    },

    formattedStringToLocalStr: function (formattedDateTime) {
        return new Date(formattedDateTime)
            .toLocaleDateString('zh-CN', {
                year: 'numeric',
                month: '2-digit',
                day: '2-digit',
                hour: '2-digit',
                minute: '2-digit',
                second: '2-digit',
            });
    },

    /**
     * 格式化的日期时间字符串 (yyyy-MM-dd HH:mm:ss) 转换为秒级时间戳
     * @param {string} formattedDateTime 格式化的日期时间字符串
     * @returns {number} 秒级时间戳
     */
    formattedStringToTimestampSeconds: function (formattedDateTime) {
        const date = this.formattedStringToDate(formattedDateTime);
        return this.dateToTimestampSeconds(date);
    },

    /**
     * 格式化的日期时间字符串 (yyyy-MM-dd HH:mm:ss) 转换为毫秒级时间戳
     * @param {string} formattedDateTime 格式化的日期时间字符串
     * @returns {number} 毫秒级时间戳
     */
    formattedStringToTimestampMilliseconds: function (formattedDateTime) {
        const date = this.formattedStringToDate(formattedDateTime);
        return this.dateToTimestampMilliseconds(date);
    },

    /**
     * 秒级时间戳转换为格式化的日期时间字符串 (yyyy-MM-dd HH:mm:ss)
     * @param {number} timestamp 秒级时间戳
     * @returns {string} 格式化的日期时间字符串
     */
    timestampSecondsToFormattedDateTime: function (timestamp) {
        const date = this.timestampSecondsToDate(timestamp);
        return this.dateToFormattedDateTime(date);
    },

    /**
     * 毫秒级时间戳转换为格式化的日期时间字符串 (yyyy-MM-dd HH:mm:ss)
     * @param {number} timestamp 毫秒级时间戳
     * @returns {string} 格式化的日期时间字符串
     */
    timestampMillisecondsToFormattedDateTime: function (timestamp) {
        const date = this.timestampMillisecondsToDate(timestamp);
        return this.dateToFormattedDateTime(date);
    },
    // 获取周几
    getWeekday() {
        const weekdays = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'];
        return weekdays[new Date().getDay()];
    },
};

export default TimeUtil;  // 或 module.exports = TimeUtil;
