package SparrowCache



type SCache struct {
    
    ht   *SHashTable
    
    
}



func NewSCache() *SCache {
    
    this := &SCache{ht:NewHashTable()}
    
    this.ht.InitHashTable()
    
    
    return this
}




func (this *SCache) Set(key,value string) error {
    
    return this.ht.Set(key,value)
    
}


func (this *SCache) Get(key string) (string,bool) {
    
    
    value,ok:= this.ht.Get(key)
    
    if !ok {
        return "",false
    }
    
    return value,ok
    
}

