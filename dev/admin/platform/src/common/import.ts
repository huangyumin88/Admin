//这个函数必须单独一个文件。否则调用过该方法的文件，再在当前这个文件内导入时会报错
//例如下面导入i18n会报：Uncaught ReferenceError: Cannot access 'batchImport' before initialization。
//报错原因：创建i18n时使用过batchImport方法，i18n会加载这个文件，而这里导入i18n，i1n8创建不成功，就会报错
//import i18n from '@/i18n'
/**
 * 批量导入
 * @param rawImportList 导入列表。（调用import.meta.globEager或import.meta.glob方法返回的数据）
 * @param level 命名层次。特别注意：如果有不同层次文件时，默认以最浅层文件为基准开始命名。正数则表示以父级文件夹（增加相应层数）为基准开始命名；负数则表示以子级文件（减去相应层数）为基准开始命名，意味着将有部分文件不会返回。例如：-1，则最浅层文件将不返回。
 * @param type  类型，默认0。0：一维对象（键名保持原样）；1：一维对象（键名小驼峰法）；2：一维对象（键名大驼峰法）；10：多维对象（键名保持原样）；11：多维对象（键名小驼峰法）；12：多维对象（键名大驼峰法）；
 * @returns 
 */
export const batchImport = async (rawImportList: any, level: number = 0, type: number = 0): Promise<{ [propName: string]: any }> => {
    let importList: { [propName: string]: any } = {}
    let keyArr: string[] = []
    let keyList: string[][] = []
    let importArr: any[] = []
    let levelOfMin: number = 0
    let importOne: { [propName: string]: any } = {}
    for (const path in rawImportList) {
        //keyArr = path.slice(0, path.lastIndexOf('.')).split('/')  //有时.不在最后
        keyArr = path.split('/')
        keyArr[keyArr.length - 1] = keyArr[keyArr.length - 1].slice(0, keyArr[keyArr.length - 1].indexOf('.'))
        keyList.push(keyArr)
        if (typeof rawImportList[path] === 'function') {
            importOne = await rawImportList[path]()
        } else {
            importOne = await rawImportList[path]
        }
        if (importOne.default) {  //有默认值只返回默认值
            importArr.push(importOne.default)
        } else {
            importArr.push(importOne)
        }
        if (keyArr.length < levelOfMin || levelOfMin == 0) {
            levelOfMin = keyArr.length
        }
    }
    const start: number = levelOfMin - level - 1 < 0 ? 0 : levelOfMin - level - 1;    //键名开始的位置
    switch (type) {
        case 0:
            for (const key in keyList) {
                const keyFinal = keyList[key].slice(start).reduce((keyFinalTmp, value) => {
                    return keyFinalTmp + value
                })
                importList[keyFinal] = importArr[key]
            }
            break;
        case 1:
            for (const key in keyList) {
                const keyFinal = keyList[key].slice(start).reduce((keyFinalTmp, value, index) => {
                    if (index == 0) {
                        return keyFinalTmp += value.split(/[\s-_]/).reduce((keyFinalTmp, value, index) => {
                            if (index == 0) {
                                return keyFinalTmp += value.slice(0, 1).toLowerCase() + value.slice(1)
                            }
                            return keyFinalTmp += value.slice(0, 1).toUpperCase() + value.slice(1)
                        }, '')
                    }
                    return keyFinalTmp += value.split(/[\s-_]/).reduce((keyFinalTmp, value) => {
                        return keyFinalTmp + value.slice(0, 1).toUpperCase() + value.slice(1)
                    }, '')
                }, '')
                importList[keyFinal] = importArr[key]
            }
            break;
        case 2:
            for (const key in keyList) {
                const keyFinal = keyList[key].slice(start).reduce((keyFinalTmp, value, index) => {
                    return keyFinalTmp += value.split(/[\s-_]/).reduce((keyFinalTmp, value) => {
                        return keyFinalTmp + value.slice(0, 1).toUpperCase() + value.slice(1)
                    }, '')
                }, '')
                importList[keyFinal] = importArr[key]
            }
            break;
        case 10:
            for (const key in keyList) {
                keyList[key].slice(start).reduce((importTmp, value, index, arr) => {
                    const keyFinal = value

                    if (index == arr.length - 1) {
                        importTmp[keyFinal] = importArr[key]
                    } else {
                        if (!(keyFinal in importTmp)) {
                            importTmp[keyFinal] = {}
                        }
                    }
                    return importTmp[keyFinal]
                }, importList)  //将importList作为importTmp的初始值，当importTmp改变，importList同时也会改变（js对象除非深复制，否则不管多少个变量都是指向同一个内存地址）
            }
            break;
        case 11:
            for (const key in keyList) {
                keyList[key].slice(start).reduce((importTmp, value, index, arr) => {
                    const keyFinal = value.split(/[\s-_]/).reduce((keyFinalTmp, value, index) => {
                        if (index == 0) {
                            return keyFinalTmp += value.slice(0, 1).toLowerCase() + value.slice(1)
                        }
                        return keyFinalTmp += value.slice(0, 1).toUpperCase() + value.slice(1)
                    }, '')

                    if (index == arr.length - 1) {
                        importTmp[keyFinal] = importArr[key]
                    } else {
                        if (!(keyFinal in importTmp)) {
                            importTmp[keyFinal] = {}
                        }
                    }
                    return importTmp[keyFinal]
                }, importList)  //将importList作为importTmp的初始值，当importTmp改变，importList同时也会改变（js对象除非深复制，否则不管多少个变量都是指向同一个内存地址）
            }
            break;
        case 12:
            for (const key in keyList) {
                keyList[key].slice(start).reduce((importTmp, value, index, arr) => {
                    const keyFinal = value.split(/[\s-_]/).reduce((keyFinalTmp, value, index) => {
                        return keyFinalTmp += value.slice(0, 1).toUpperCase() + value.slice(1)
                    }, '')

                    if (index == arr.length - 1) {
                        importTmp[keyFinal] = importArr[key]
                    } else {
                        if (!(keyFinal in importTmp)) {
                            importTmp[keyFinal] = {}
                        }
                    }
                    return importTmp[keyFinal]
                }, importList)  //将importList作为importTmp的初始值，当importTmp改变，importList同时也会改变（js对象除非深复制，否则不管多少个变量都是指向同一个内存地址）
            }
            break;
    }
    return importList
}
/*--------使用方式 开始--------*/
// console.log(await batchImport(import.meta.globEager('@/i18n/language/**/*.ts'), 1, 10))
// console.log(await batchImport(import.meta.globEager('@/api/**/*.ts')))
// console.log(await batchImport(import.meta.globEager('@/../node_modules/element-plus/dist/locale/*.min.mjs')))
/*--------使用方式 结束--------*/