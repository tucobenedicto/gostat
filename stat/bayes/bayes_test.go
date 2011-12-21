/*
R
> library("Bolstad")
> binobp(x, n, a = 1, b = 1, ret = FALSE)
Error in binobp(x, n, a = 1, b = 1, ret = FALSE) : object 'x' not found
> binobp(x=10, n=20, a = 1, b = 1, ret = FALSE)
Posterior Mean           :  0.5 
Posterior Variance       :  0.0108696 
Posterior Std. Deviation :  0.1042572 

Prob.	Quantile 
------	---------
0.005	0.2424494
0.01	0.2642106
0.025	0.2978068
0.05	0.3281087
0.5	0.5
0.95	0.6718913
0.975	0.7021932
0.99	0.7357894
0.995	0.7575506
> 
*/


package bayes

import (
	"testing"
	"fmt"
//	s "gostat.googlecode.com/hg/stat"
)

// test against known values
func TestBinomBetaPriQtl(t *testing.T) {
	var (
		α, β, y float64
		i, k, n int64
	)


	// edit the following values:  >>>

k=10
n=20
α=1
β=1
	p:=[]float64{0.005,0.01,0.025,0.05,0.5,0.95,0.975,0.99,0.995}
	x:=[]float64{0.2424494,0.2642106,0.2978068,0.3281087,0.5,0.6718913,0.7021932,0.7357894,0.7575506}


	// <<<

	fmt.Println("test for BinomBetaPriQtl")
	for i = 0; i < int64(len(p)); i++ {
		y=BinomPi_Qtl_BPri(k, n , α, β, p[i])
			if !check(x[i], y){
				t.Error()
				fmt.Println(p[i], x[i], y)

			}
		y=BinomPiPostMean(α, β, n, k)
			if !check(y, 0.5){
				t.Error()
				fmt.Println("posterior mean failed")

			}

	}
}


// postPDF_ FlatPri (8, 6){2.514963e-16 1.601541e-14 1.818782e-13 1.019354e-12 3.879562e-12 1.155875e-11 2.908412e-11 6.466759e-11 1.308249e-10 2.456577e-10 4.342959e-10 7.304989e-10 1.178411e-09 1.834471e-09 2.769474e-09 4.070778e-09 5.84466e-09 8.218812e-09 1.134498e-08 1.540176e-08 2.059748e-08 2.717328e-08 3.540627e-08 4.561286e-08 5.815219e-08 7.342968e-08 9.190076e-08 1.140746e-07 1.405182e-07 1.718601e-07 2.087948e-07 2.520873e-07 3.025766e-07 3.611812e-07 4.28903e-07 5.068321e-07 5.961519e-07 6.981437e-07 8.141917e-07 9.457883e-07 1.094539e-06 1.262168e-06 1.450522e-06 1.661580e-06 1.897452e-06 2.160391e-06 2.452796e-06 2.777215e-06 3.136358e-06 3.533094e-06 3.970465e-06 4.451687e-06 4.980158e-06 5.559463e-06 6.193381e-06 6.885894e-06 7.641188e-06 8.463662e-06 9.357936e-06 1.032886e-05 1.13815e-05 1.252118e-05 1.375347e-05 1.508418e-05 1.651939e-05 1.806543e-05 1.972893e-05 2.151677e-05 2.343614e-05 2.549452e-05 2.769968e-05 3.00597e-05 3.258298e-05 3.527825e-05 3.815455e-05 4.122127e-05 4.448813e-05 4.796522e-05 5.166298e-05 5.559219e-05 5.976404e-05 6.419009e-05 6.888225e-05 7.385288e-05 7.91147e-05 8.468085e-05 9.05649e-05 9.67808e-05 0.000103343 0.0001102663 0.0001175660 0.0001252579 0.0001333581 0.0001418834 0.0001508508 0.0001602779 0.000170183 0.0001805845 0.0001915017 0.0002029540 0.0002149617 0.0002275454 0.0002407262 0.0002545259 0.0002689668 0.0002840717 0.0002998638 0.0003163673 0.0003336066 0.0003516068 0.0003703935 0.0003899931 0.0004104325 0.000431739 0.0004539408 0.0004770666 0.0005011457 0.0005262082 0.0005522845 0.0005794061 0.0006076048 0.000636913 0.0006673643 0.0006989923 0.0007318317 0.0007659177 0.0008012864 0.0008379744 0.000876019 0.0009154582 0.000956331 0.0009986768 0.001042536 0.001087949 0.001134958 0.001183605 0.001233934 0.001285988 0.001339813 0.001395453 0.001452955 0.001512366 0.001573733 0.001637106 0.001702533 0.001770065 0.001839753 0.001911649 0.001985804 0.002062273 0.002141109 0.002222369 0.002306107 0.00239238 0.002481246 0.002572763 0.002666991 0.002763989 0.002863818 0.002966540 0.003072219 0.003180916 0.003292697 0.003407627 0.003525771 0.003647197 0.003771974 0.003900168 0.00403185 0.004167091 0.004305962 0.004448535 0.004594883 0.004745081 0.004899202 0.005057325 0.005219523 0.005385877 0.005556464 0.005731363 0.005910656 0.006094423 0.006282746 0.006475709 0.006673396 0.006875892 0.007083282 0.007295653 0.007513094 0.007735692 0.007963537 0.00819672 0.008435333 0.008679466 0.008929214 0.009184671 0.009445932 0.009713093 0.00998625 0.01026550 0.01055095 0.01084268 0.01114082 0.01144544 0.01175667 0.01207459 0.01239932 0.01273096 0.01306961 0.01341539 0.01376840 0.01412875 0.01449654 0.0148719 0.01525493 0.01564574 0.01604444 0.01645116 0.01686601 0.01728909 0.01772053 0.01816045 0.01860897 0.01906619 0.01953225 0.02000727 0.02049136 0.02098465 0.02148726 0.02199932 0.02252095 0.02305229 0.02359344 0.02414455 0.02470575 0.02527715 0.02585890 0.02645111 0.02705394 0.02766749 0.02829192 0.02892735 0.02957393 0.03023177 0.03090103 0.03158183 0.03227432 0.03297864 0.03369492 0.03442330 0.03516393 0.03591695 0.03668249 0.03746071 0.03825174 0.03905573 0.03987282 0.04070317 0.04154690 0.04240419 0.04327516 0.04415997 0.04505877 0.04597171 0.04689893 0.04784059 0.04879684 0.04976783 0.05075371 0.05175463 0.05277076 0.05380223 0.05484922 0.05591186 0.05699032 0.05808476 0.05919532 0.06032218 0.06146547 0.06262537 0.06380202 0.0649956 0.06620625 0.06743413 0.06867942 0.06994226 0.07122282 0.07252125 0.07383773 0.0751724 0.07652544 0.077897 0.07928726 0.08069636 0.08212448 0.08357177 0.0850384 0.08652453 0.08803033 0.08955597 0.0911016 0.09266739 0.0942535 0.09586011 0.09748737 0.09913545 0.1008045 0.1024947 0.1042063 0.1059393 0.1076939 0.1094704 0.1112688 0.1130894 0.1149323 0.1167977 0.1186857 0.1205965 0.1225303 0.1244872 0.1264674 0.1284710 0.1304983 0.1325494 0.1346244 0.1367236 0.1388470 0.1409949 0.1431674 0.1453647 0.1475869 0.1498342 0.1521068 0.1544047 0.1567283 0.1590776 0.1614528 0.1638541 0.1662816 0.1687355 0.1712159 0.1737231 0.1762570 0.1788180 0.1814062 0.1840217 0.1866647 0.1893353 0.1920338 0.1947601 0.1975146 0.2002973 0.2031085 0.2059481 0.2088165 0.2117138 0.2146400 0.2175954 0.2205801 0.2235943 0.2266380 0.2297115 0.2328149 0.2359482 0.2391118 0.2423056 0.2455299 0.2487848 0.2520704 0.2553868 0.2587343 0.2621129 0.2655227 0.2689640 0.2724367 0.2759411 0.2794774 0.2830455 0.2866456 0.290278 0.2939426 0.2976396 0.3013692 0.3051314 0.3089264 0.3127543 0.3166152 0.3205092 0.3244364 0.328397 0.3323911 0.3364187 0.34048 0.344575 0.348704 0.3528669 0.3570639 0.3612951 0.3655606 0.3698605 0.3741948 0.3785637 0.3829672 0.3874055 0.3918786 0.3963867 0.4009297 0.4055078 0.4101211 0.4147697 0.4194535 0.4241727 0.4289274 0.4337176 0.4385434 0.4434049 0.4483021 0.4532351 0.458204 0.4632087 0.4682495 0.4733262 0.478439 0.4835879 0.488773 0.4939943 0.4992519 0.5045457 0.5098759 0.5152424 0.5206453 0.5260846 0.5315604 0.5370727 0.5426215 0.5482068 0.5538286 0.559487 0.5651819 0.5709135 0.5766816 0.5824863 0.5883276 0.5942055 0.60012 0.6060711 0.6120587 0.618083 0.6241437 0.630241 0.6363749 0.6425452 0.648752 0.6549952 0.6612749 0.6675909 0.6739433 0.680332 0.6867569 0.693218 0.6997153 0.7062487 0.7128182 0.7194237 0.7260651 0.7327424 0.7394555 0.7462043 0.7529889 0.759809 0.7666646 0.7735556 0.780482 0.7874436 0.7944404 0.8014723 0.808539 0.8156407 0.8227772 0.8299482 0.8371538 0.8443939 0.8516682 0.8589767 0.8663192 0.8736956 0.8811059 0.8885497 0.8960271 0.9035378 0.9110818 0.9186588 0.9262687 0.9339113 0.9415866 0.9492942 0.9570341 0.964806 0.97261 0.9804455 0.9883126 0.996211 1.004141 1.012101 1.020093 1.028114 1.036167 1.044249 1.052361 1.060503 1.068675 1.076875 1.085105 1.093363 1.101650 1.109966 1.118309 1.126680 1.135079 1.143505 1.151958 1.160438 1.168944 1.177477 1.186035 1.194619 1.203229 1.211863 1.220523 1.229206 1.237914 1.246646 1.255402 1.264180 1.272982 1.281806 1.290653 1.299521 1.308411 1.317323 1.326255 1.335208 1.344181 1.353174 1.362186 1.371218 1.380268 1.389337 1.398423 1.407528 1.416649 1.425788 1.434943 1.444113 1.4533 1.462502 1.471719 1.480950 1.490195 1.499454 1.508725 1.51801 1.527307 1.536615 1.545935 1.555266 1.564608 1.573959 1.58332 1.592690 1.602069 1.611456 1.620850 1.630252 1.639660 1.649075 1.658495 1.667921 1.677351 1.686785 1.696223 1.705664 1.715108 1.724554 1.734001 1.74345 1.752899 1.762348 1.771796 1.781243 1.790689 1.800132 1.809573 1.81901 1.828443 1.837872 1.847295 1.856713 1.866124 1.875529 1.884926 1.894315 1.903696 1.913067 1.922428 1.931779 1.941119 1.950447 1.959762 1.969065 1.978354 1.987629 1.996889 2.006133 2.015361 2.024572 2.033766 2.042941 2.052098 2.061235 2.070353 2.079449 2.088524 2.097576 2.106606 2.115612 2.124594 2.133551 2.142482 2.151387 2.160265 2.169116 2.177938 2.18673 2.195493 2.204225 2.212926 2.221595 2.230231 2.238833 2.247402 2.255935 2.264433 2.272894 2.281318 2.289704 2.298051 2.306359 2.314627 2.322854 2.331039 2.339182 2.347281 2.355337 2.363347 2.371312 2.379231 2.387103 2.394927 2.402702 2.410428 2.418103 2.425728 2.433301 2.440821 2.448288 2.455701 2.463059 2.470361 2.477607 2.484795 2.491925 2.498996 2.506008 2.512959 2.519849 2.526676 2.533441 2.540142 2.546778 2.553349 2.559854 2.566292 2.572662 2.578963 2.585196 2.591357 2.597448 2.603467 2.609414 2.615287 2.621085 2.626809 2.632456 2.638027 2.643521 2.648936 2.654272 2.659528 2.664703 2.669797 2.674809 2.679737 2.684581 2.689341 2.694014 2.698602 2.703102 2.707515 2.711838 2.716072 2.720216 2.724268 2.728229 2.732097 2.735871 2.739551 2.743136 2.746625 2.750017 2.753312 2.756509 2.759606 2.762604 2.765502 2.768298 2.770993 2.773584 2.776072 2.778456 2.780735 2.782908 2.784975 2.786934 2.788786 2.790529 2.792162 2.793686 2.795098 2.796400 2.797589 2.798665 2.799627 2.800476 2.801209 2.801827 2.802328 2.802713 2.802980 2.803128 2.803158 2.803069 2.802859 2.802528 2.802076 2.801502 2.800806 2.799986 2.799042 2.797974 2.796781 2.795463 2.794018 2.792447 2.790749 2.788923 2.786968 2.784885 2.782673 2.780331 2.777859 2.775256 2.772521 2.769655 2.766657 2.763526 2.760263 2.756866 2.753335 2.749669 2.745869 2.741935 2.737864 2.733658 2.729316 2.724837 2.720222 2.715470 2.710580 2.705552 2.700387 2.695083 2.689641 2.68406 2.67834 2.672481 2.666483 2.660345 2.654068 2.647651 2.641093 2.634396 2.627558 2.62058 2.613462 2.606203 2.598803 2.591263 2.583583 2.575762 2.5678 2.559697 2.551455 2.543071 2.534548 2.525884 2.517080 2.508135 2.499051 2.489828 2.480464 2.470962 2.46132 2.451540 2.441621 2.431563 2.421368 2.411035 2.400564 2.389957 2.379213 2.368333 2.357317 2.346165 2.334879 2.323459 2.311905 2.300217 2.288397 2.276444 2.26436 2.252145 2.2398 2.227325 2.214722 2.20199 2.189131 2.176145 2.163033 2.149796 2.136435 2.122951 2.109345 2.095617 2.081769 2.067801 2.053714 2.039511 2.025190 2.010755 1.996206 1.981543 1.966769 1.951884 1.936890 1.921788 1.906579 1.891264 1.875846 1.860325 1.844703 1.828981 1.813160 1.797243 1.781231 1.765126 1.748928 1.732640 1.716264 1.699801 1.683252 1.666621 1.649907 1.633115 1.616245 1.599299 1.582279 1.565187 1.548026 1.530798 1.513504 1.496146 1.478728 1.461251 1.443717 1.426129 1.408490 1.390801 1.373065 1.355284 1.337462 1.3196 1.301702 1.283769 1.265805 1.247813 1.229795 1.211753 1.193692 1.175614 1.157521 1.139417 1.121305 1.103188 1.085069 1.066952 1.048839 1.030734 1.012640 0.9945614 0.9765006 0.9584613 0.9404472 0.9224619 0.904509 0.8865922 0.8687153 0.850882 0.8330963 0.815362 0.7976832 0.7800638 0.7625078 0.7450196 0.7276031 0.7102627 0.6930027 0.6758273 0.6587411 0.6417485 0.624854 0.6080622 0.5913777 0.5748053 0.5583496 0.5420156 0.5258081 0.5097321 0.4937925 0.4779944 0.462343 0.4468434 0.431501 0.4163209 0.4013087 0.3864697 0.3718095 0.3573337 0.3430478 0.3289577 0.3150691 0.3013878 0.2879198 0.2746711 0.2616476 0.2488556 0.2363012 0.2239908 0.2119305 0.2001270 0.1885865 0.1773157 0.1663213 0.1556099 0.1451883 0.1350633 0.1252420 0.1157312 0.1065381 0.09766977 0.08913356 0.08093672 0.07308667 0.06559086 0.05845685 0.05169228 0.04530486 0.03930239 0.03369275 0.02848389 0.02368387 0.01930080 0.01534292 0.0118185 0.00873594 0.006103708 0.003930358 0.002224536 0.0009949737 0.0002504918}
