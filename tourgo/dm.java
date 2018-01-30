int alphaBeta(int Alpha ,int beta,int depth) {
    //time out
    if(dataTime.Now.Ticks-starttime >=object.limttime)
    return -9999;
    if ((depth ==0) || (checkposition==true)) // do sau bang ko or la win : 
    return evaluate(depth);// ham danh gia evaluate(position)
    else 
    {
int best,value;
best =-1000// best =-infinitive
arraylist totalPosition= new arraylist();
totalPosition = generateMove(Positionnow);
while (i <totalPosition.size && best <beta)
(
    if (best>Alpha) Alpha=best;
    domove(a[i])
    value=-alphaBeta(-best,-Alpha,depth-1);
    remove(a[i]);
    if((math.abs.(value)!=1000)&&value>best)
    {
        best=value;
        
    }
)
return best;
    }
}
