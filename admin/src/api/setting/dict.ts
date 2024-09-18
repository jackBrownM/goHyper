import request from '@/utils/request'

// 字典数据列表
export function dictDataAll(params: any) {
    return request.get(
        { url: '/setting/dict/data/all', params },
        {
            ignoreCancelToken: true
        }
    )
}